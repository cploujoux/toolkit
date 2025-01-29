# Import necessary modules
import functools
import inspect
from logging import getLogger

from langgraph.checkpoint.memory import MemorySaver
from langgraph.prebuilt import create_react_agent

from beamlit.api.models import get_model, list_models
from beamlit.authentication import new_client
from beamlit.common.settings import init
from beamlit.errors import UnexpectedStatus
from beamlit.functions import get_functions
from beamlit.models import Agent, AgentMetadata, AgentSpec

from .chat import get_chat_model_full


def agent(
    agent: Agent | dict = None,
    override_model=None,
    override_agent=None,
    override_functions=None,
    remote_functions=None,
):
    logger = getLogger(__name__)
    try:
        if agent is not None and not isinstance(agent, dict):
            raise Exception(
                'agent must be a dictionary, example: @agent(agent={"metadata": {"name": "my_agent"}})'
            )

        client = new_client()
        chat_model = override_model or None
        settings = init()

        def wrapper(func):
            agent_kwargs = any(
                param.name == "agent"
                for param in inspect.signature(func).parameters.values()
            )
            model_kwargs = any(
                param.name == "model"
                for param in inspect.signature(func).parameters.values()
            )
            functions_kwargs = any(
                param.name == "functions"
                for param in inspect.signature(func).parameters.values()
            )
            @functools.wraps(func)
            def wrapped(*args, **kwargs):
                if agent_kwargs:
                    kwargs["agent"] = settings.agent.agent
                if model_kwargs:
                    kwargs["model"] = settings.agent.chat_model
                if functions_kwargs:
                    kwargs["functions"] = settings.agent.functions
                return func(*args, **kwargs)

            return wrapped

        if agent is not None:
            metadata = AgentMetadata(**agent.get("metadata", {}))
            spec = AgentSpec(**agent.get("spec", {}))
            agent = Agent(metadata=metadata, spec=spec)
            if agent.spec.model and chat_model is None:
                try:
                    response = get_model.sync_detailed(
                        agent.spec.model, environment=settings.environment, client=client
                    )
                    settings.agent.model = response.parsed
                except UnexpectedStatus as e:
                    if e.status_code == 404 and settings.environment != "production":
                        try:
                            response = get_model.sync_detailed(
                                agent.spec.model, environment="production", client=client
                            )
                            settings.agent.model = response.parsed
                        except UnexpectedStatus as e:
                            if e.status_code == 404:
                                raise ValueError(f"Model {agent.spec.model} not found")
                    else:
                        raise e
                except Exception as e:
                    raise e

                if settings.agent.model:
                    chat_model, provider, model = get_chat_model_full(agent.spec.model, settings.agent.model)
                    settings.agent.chat_model = chat_model
                    logger.info(f"Chat model configured, using: {provider}:{model}")

        if override_functions is not None:
            functions = override_functions
        else:
            functions = get_functions(
                client=client,
                dir=settings.agent.functions_directory,
                remote_functions=remote_functions,
                chain=agent.spec.agent_chain,
            remote_functions_empty=not remote_functions,
            warning=chat_model is not None,
        )

        settings.agent.functions = functions

        if override_agent is None:
            if chat_model is None:
                models_select = ""
                try:
                    models = list_models.sync_detailed(
                        environment=settings.environment, client=client
                    )
                    models = ", ".join([model.metadata.name for model in models.parsed])
                    models_select = f"You can select one from the your models: {models}"
                except Exception:
                    pass

                raise ValueError(f"You must provide a model.\n"
                    f"{models_select}\n"
                    f"You can create one at {settings.app_url}/{settings.workspace}/global-inference-network/models/create\n"
                    "Add it to your agent spec\n"
                    "agent={\n"
                    "    \"metadata\": {\n"
                    f"        \"name\": \"{agent.metadata.name}\",\n"
                    "    },\n"
                    "    \"spec\": {\n"
                    "        \"model\": \"MODEL_NAME\",\n"
                    f"        \"description\": \"{agent.spec.description}\",\n"
                    "    },\n"
                    "}")
            memory = MemorySaver()
            if len(functions) == 0:
                raise ValueError("You can define this function in directory "
                    f'"{settings.agent.functions_directory}". Here is a sample function you can use:\n\n'
                    "from beamlit.functions import function\n\n"
                    "@function()\n"
                    "def hello_world(query: str):\n"
                    "    return 'Hello, world!'\n")
            _agent = create_react_agent(chat_model, functions, checkpointer=memory)
            settings.agent.agent = _agent
        else:
            settings.agent.agent = override_agent
        return wrapper
    except Exception as e:
        logger.error(f"Error in agent decorator: {e!s} at line {e.__traceback__.tb_lineno}")
        raise e

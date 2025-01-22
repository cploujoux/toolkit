# Import necessary modules
import functools
from logging import getLogger

from langgraph.checkpoint.memory import MemorySaver
from langgraph.prebuilt import create_react_agent

from beamlit.api.models import get_model
from beamlit.authentication import new_client
from beamlit.common.settings import init
from beamlit.errors import UnexpectedStatus
from beamlit.functions import get_functions
from beamlit.models import Agent, AgentMetadata, AgentSpec

from .chat import get_chat_model


def agent(
    agent: Agent | dict = None,
    override_model=None,
    override_agent=None,
    mcp_hub=None,
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
            @functools.wraps(func)
            def wrapped(*args, **kwargs):
                return func(
                    settings.agent.agent,
                    settings.agent.chat_model,
                    settings.agent.functions,
                    *args,
                    **kwargs,
                )

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
                    chat_model, provider, model = get_chat_model(agent.spec.model, settings.agent.model)
                    settings.agent.chat_model = chat_model
                    logger.info(f"Chat model configured, using: {provider}:{model}")

        functions = get_functions(
            client=client,
            dir=settings.agent.functions_directory,
            mcp_hub=mcp_hub,
            remote_functions=remote_functions,
            chain=agent.spec.agent_chain,
            remote_functions_empty=not remote_functions,
        )
        settings.agent.functions = functions
        
        if override_agent is None and len(functions) == 0:
            raise ValueError(
                "You must define at least one function, you can define this function in directory "
                f'"{settings.agent.functions_directory}". Here is a sample function you can use:\n\n'
                "from beamlit.functions import function\n\n"
                "@function()\n"
                "def hello_world(query: str):\n"
                "    return 'Hello, world!'\n"
            )

        if override_agent is None and chat_model is not None:
            memory = MemorySaver()
            _agent = create_react_agent(chat_model, functions, checkpointer=memory)
            settings.agent.agent = _agent
        else:
            settings.agent.agent = override_agent
        return wrapper
    except Exception as e:
        logger.error(f"Error in agent decorator: {e!s} at line {e.__traceback__.tb_lineno}")
        raise e

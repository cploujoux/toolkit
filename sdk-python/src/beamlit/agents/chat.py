from logging import getLogger
from typing import Tuple, Union

from langchain_core.language_models import BaseChatModel

from beamlit.api.models import get_model
from beamlit.authentication import get_authentication_headers, new_client
from beamlit.common.settings import get_settings
from beamlit.models import Model

logger = getLogger(__name__)


def get_base_url(name: str):
    settings = get_settings()
    return f"{settings.run_url}/{settings.workspace}/models/{name}/v1"


def get_mistral_chat_model(**kwargs):
    from langchain_mistralai.chat_models import ChatMistralAI  # type: ignore

    return ChatMistralAI(**kwargs)


def get_openai_chat_model(**kwargs):
    from langchain_openai import ChatOpenAI  # type: ignore

    return ChatOpenAI(**kwargs)


def get_anthropic_chat_model(**kwargs):
    from langchain_anthropic import ChatAnthropic  # type: ignore

    return ChatAnthropic(**kwargs)

def get_xai_chat_model(**kwargs):
    from langchain_xai import ChatXAI  # type: ignore

    return ChatXAI(**kwargs)

def get_cohere_chat_model(**kwargs):
    from langchain_cohere import ChatCohere  # type: ignore

    return ChatCohere(**kwargs)

def get_deepseek_chat_model(**kwargs):
    from langchain_deepseek import ChatDeepSeek  # type: ignore

    return ChatDeepSeek(**kwargs)

def get_chat_model(name: str, agent_model: Union[Model, None] = None) -> BaseChatModel:
    [chat_model, _, __] = get_chat_model_full(name, agent_model)
    return chat_model

def get_chat_model_full(name: str, agent_model: Union[Model, None] = None) -> Tuple[BaseChatModel, str, str]:
    settings = get_settings()
    client = new_client()

    if agent_model is None:
        try:
            agent_model = get_model.sync(name, client=client, environment=settings.environment)
        except Exception:
            logger.warning(f"Model {name} not found, defaulting to gpt-4o-mini")

    environment = (agent_model and agent_model.metadata and agent_model.metadata.environment) or settings.environment
    headers = get_authentication_headers(settings)
    headers["X-Beamlit-Environment"] = environment

    jwt = headers.get("X-Beamlit-Authorization", "").replace("Bearer ", "")
    params = {"environment": environment}
    chat_classes = {
        "openai": {
            "func": get_openai_chat_model,
            "kwargs": {
                "http_async_client": client.get_async_httpx_client(),
                "http_client": client.get_httpx_client(),
            },
        },
        "anthropic": {
            "func": get_anthropic_chat_model,
            "kwargs": {
                "base_url": get_base_url(name).replace("/v1", ""),
            },
            "remove_kwargs": ["default_query"]
        },
        "mistral": {
            "func": get_mistral_chat_model,
            "kwargs": {
                "api_key": jwt,
            },
        },
        "xai": {
            "func": get_xai_chat_model,
            "kwargs": {
                "api_key": jwt,
                "xai_api_base": get_base_url(name),
            },
            "remove_kwargs": ["base_url"],
        },
        "cohere": {
            "func": get_cohere_chat_model,
            "kwargs": {
                "cohere_api_key": jwt,
                "base_url": get_base_url(name).replace("/v1", ""),
            },
        },
        "deepseek": {
            "func": get_deepseek_chat_model,
            "kwargs": {
                "api_key": jwt,
            },
        },
    }

    provider = (
        agent_model
        and agent_model.spec
        and agent_model.spec.runtime
        and agent_model.spec.runtime.type_
    )
    if not provider:
        logger.warning("Provider not found in agent model, defaulting to OpenAI")
        provider = "openai"

    model = (
        agent_model
        and agent_model.spec
        and agent_model.spec.runtime
        and agent_model.spec.runtime.model
    )
    if not model:
        logger.warning("Model not found in agent model, defaulting to gpt-4o-mini")
        model = "gpt-4o-mini"

    kwargs = {
        "model": model,
        "base_url": get_base_url(name),
        "default_query": params,
        "default_headers": headers,
        "api_key": "fake_api_key",
        "temperature": 0,
    }
    chat_class = chat_classes.get(provider)
    if not chat_class:
        logger.warning(f"Provider {provider} not currently supported, defaulting to OpenAI")
        chat_class = chat_classes["openai"]
    if "kwargs" in chat_class:
        kwargs.update(chat_class["kwargs"])
    if "remove_kwargs" in chat_class:
        for key in chat_class["remove_kwargs"]:
            kwargs.pop(key, None)
    return chat_class["func"](**kwargs), provider, model

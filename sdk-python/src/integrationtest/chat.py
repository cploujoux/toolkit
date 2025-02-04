import asyncio

from langgraph.graph.graph import CompiledGraph

from beamlit.agents import agent, get_chat_model
from beamlit.common import init
from beamlit.functions import function

settings = init()
chat = get_chat_model("gpt-4o-mini") # all good
# chat = get_chat_model("ministral-3b-2410") # all good
# chat = get_chat_model("cohere-command-r-plus") # all good, astream does not work
# chat = get_chat_model("claude-3-5-sonnet") # all good
# chat = get_chat_model("deepseek-chat")
# chat = get_chat_model("xai-grok-beta") # all good

print(chat)
@function()
async def get_weather(location: str):
    """Get the weather for a given location"""
    print("Using getWeather");
    return "The weather in " + location + " is sunny"

@agent(
    override_model=chat,
    override_functions=[get_weather],
)
async def main(request, agent: CompiledGraph):
    config = {"configurable": {"thread_id": "thread_id"}}
    agent_body = {"messages": [("user", "What's the weather in san francisco ?")]}
    response = None
    for chunk in agent.stream(agent_body, config=config):
        response = chunk
    print(response)


asyncio.run(main({}))

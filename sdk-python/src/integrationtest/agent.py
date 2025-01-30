import asyncio
import uuid

from beamlit.agents import agent
from beamlit.common import init

settings = init()

@agent(
    agent={
        "metadata": {
            "name": "agent-custom",
            "environment": "production",
        },
        "spec": {
            "description": "A chat agent using Beamlit to handle your tasks.",
            "model": "gpt-4o-mini",
        },
    },
    remote_functions=["brave-search"],
)
async def main(
    input, agent,
):
    agent_config = {"configurable": {"thread_id": str(uuid.uuid4())}}


    agent_body = {"messages": [("user", input)]}
    responses = []

    async for chunk in agent.astream(agent_body, config=agent_config):
        responses.append(chunk)
    content = responses[-1]
    return content["agent"]["messages"][-1].content


if __name__ == "__main__":
    async def check():
        input = "What is the weather in Paris ?"
        response = await main(input)
        print(response)

    asyncio.run(check())

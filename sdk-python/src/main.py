from beamlit.api.functions import get_function
from beamlit.authentication import (RunClientWithCredentials,
                                    new_client_with_credentials)
from beamlit.credentials import load_credentials

credentials = load_credentials("development")
config = RunClientWithCredentials(
    credentials=credentials,
    workspace="development",
)
client = new_client_with_credentials(config)

with client as client:
    response = get_function.sync_detailed("github", client=client)

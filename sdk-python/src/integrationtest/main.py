from beamlit.api.models import get_model, list_models
from beamlit.authentication import get_authentication_headers, new_client
from beamlit.common.settings import init
from beamlit.models import Model
from beamlit.run import RunClient
from beamlit.deploy import generate_beamlit_deployment

settings = init()
client = new_client()
models = list_models.sync(client=client)

print(models)
generate_beamlit_deployment(".beamlit")
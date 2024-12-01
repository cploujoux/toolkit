from beamlit import Client
from beamlit.api.functions import get_function
from beamlit.models import Function
from beamlit.types import Response

client = Client(base_url="https://api.beamlit.dev/v0", raise_on_unexpected_status=True)
with client as client:
    function: Function = get_function.sync("github", client=client)
    print(function)

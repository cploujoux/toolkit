from beamlit.authentication import new_client_with_credentials
from beamlit.credentials import load_credentials

credentials = load_credentials()
print(credentials)
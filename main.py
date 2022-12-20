import requests
import pandas as pd
from pyshapes import shapes

# Définissez l'URL de l'API REST de NetBox et votre jeton d'authentification
api_url = "https://demo.netbox.dev/api"
headers = {"Authorization": "Token 607021d8a79b515d8a813858e451803c026cc25a"}

# Effectuez une requête GET pour récupérer la liste des appareils
response = requests.get(f"{api_url}/dcim/devices/", headers=headers)

# Vérifiez le statut de la réponse pour vous assurer qu'elle a réussi
if response.status_code != 200:
    raise Exception(f"API request failed with status code {response.status_code}")

# Chargez les données de la réponse dans un DataFrame Pandas
devices_df = pd.DataFrame(response.json()["results"])

# Créez une liste de shapes pour Diagramme.net à partir des données du DataFrame
shapes_list = []
for _, row in devices_df.iterrows():
    shape = shapes.Rectangle(
        text=row["name"],
        width=100,
        height=50,
        fill="#ffffff",
        stroke="#000000",
        stroke_width=2,
    )
    shapes_list.append(shape)

# Enregistrez la liste de shapes dans un fichier JSON
shapes.save_shapes(shapes_list, "devices.json")

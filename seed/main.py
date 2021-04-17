import time

from gql import gql, Client
from gql.transport.aiohttp import AIOHTTPTransport
from gql.transport.exceptions import TransportQueryError
from os import walk


class Connection:
    transport = None
    client = None

    def fetch_connection(self, headers=None):
        # Select your transport with a defined url endpoint
        if headers is not None:
            self.transport = AIOHTTPTransport(url="http://api:8080/query", headers=headers)
        else:
            self.transport = AIOHTTPTransport(url="http://api:8080/query")

        # Create a GraphQL client using the defined transport
        self.client = Client(transport=self.transport, fetch_schema_from_transport=True)


connection = Connection()


def signup_as_company():
    company = {
        'name': "Lenfoirêt",
        'logo': "http://www.net-parodies.com/medias/images/62843556logo2-jpg.jpg?fx=r_250_250",
        'description': "En france, c'est la crise et il y a plus de trois millions et demi de mal logés.. mais heureusement il y a la fôret: 200 000m².",
        'tel': "0506012939",
        'user': {
            'username': "lenfoiret@lenfoiret.com",
            'password': "lenfoiret"
        }
    }
    query = gql(
        """
          mutation($input: CompanyInput!) {
            signupAsCompany(input: $input) {
              user {
                ID
              }
              token
            }
          }
    """
    )

    # Execute the query on the transport
    result = connection.client.execute(query, variable_values={
        'input': company
    })
    return result['signupAsCompany']


def create_estate_agents():
    estate_agent1 = {
        'about': "Chestnut, short hair almost fully covers a strong, friendly face. Clear blue eyes, set rooted within their sockets, watch anxiously over the city they've protected for so long. Fair skin delightfully compliments his nose and mouth and leaves a heartbreaking memory of his luck in battles. This is the face of Brooks Eustice, a true angel among werewolves. He stands tall above others, despite his brawny frame. There's something extraordinary about him, perhaps it's a feeling of indifference or perhaps it's simply his decency. But nonetheless, people tend to buy him a drink, while hoping to one day follow in his footsteps.",
        'skills': ["LISTENING", "SPANISH", "NEGOCIATION"],
        'specialities': ["LUXURY", "RESIDENTIAL"],
        'tel': "0656033923",
        'user': {
            'username': "jean@lenfoiret.com",
            'password': "jean"
        }
    }
    estate_agent2 = {
        'about': "Brown, coily hair gently hangs over a full, wild face. Dancing gray eyes, set far within their sockets, watch vigilantly over the river they've grieved with for so long. Fallen debry left a mark stretching from just under the right eye , running towards her right nostril and ending on her right cheek and leaves a lasting punishment of departed loved ones. The is the face of Zonkaja Burninghorn, a true utopian among orcs. She stands common among others, despite her big frame. There's something captivating about her, perhaps it's her sympathy or perhaps it's simply her unfortunate past. But nonetheless, people tend to pretend to be her best friend, while helping her out in any way they can.",
        'skills': ["COMMUNICATING", "FRENCH", "REMOTE_WORKING", "HARD_WORKING"],
        'specialities': ["PROPERTY_MANAGEMENT", "NEW_CONSTRUCTION", "FARMS"],
        'tel': "0674159422",
        'user': {
            'username': "jade@lenfoiret.com",
            'password': "jade"
        }
    }
    query = gql(
        """
          mutation($input: EstateAgentInput) {
            createEstateAgentUser(input: $input) {
              ID
              user {
                username
              }
            }
          }
    """
    )

    # Execute the query on the transport
    result1 = connection.client.execute(query, variable_values={
        'input': estate_agent1
    })
    result1['createEstateAgentUser']['user']['password'] = estate_agent1['user']['password']
    result2 = connection.client.execute(query, variable_values={
        'input': estate_agent2
    })
    result2['createEstateAgentUser']['user']['password'] = estate_agent2['user']['password']

    return [result1['createEstateAgentUser'], result2['createEstateAgentUser']]


def get_files_in_directory(path):
    files = []
    _, _, filenames = next(walk(path))

    for filename in filenames:
        files.append(open(path + filename, 'rb'))
    return files


def create_properties(count=1):
    properties = []

    properties1 = [
        {
            'area': 86,
            'country': 'France',
            'cityName': 'Rennes',
            'address': '36 rue Danton',
            'postalCode': '35700',
            'type': 'Appartement',
            'photos': get_files_in_directory('/app/photos/11/'),
            'model': get_files_in_directory('/app/models/11/')[0],
            'badges': ['Garden', 'Caretaker', 'Garage'],
            'description': 'RENNES LA BELLANGERAIS. Au 4ème et dernier étage d\'une résidence de 2014, dans un environnement verdoyant et sans vis-à-vis, beau T4 Duplex de 86m² environ. A proximité immédiate de la coulée verte, du CHRU La Tauvrais, et du Conseil Régional.',
            'rooms': 4,
            'bedrooms': 2,
            'furnished': True,
            'energyRating': 'B',
            'rentAmount': 1072,
            'chargesAmount': 200
        },
        {
            'area': 85,
            'country': 'France',
            'cityName': 'La Ferté-Macé',
            'address': '9222 Rue des Tisserands',
            'postalCode': '61600',
            'type': 'Appartement',
            'photos': get_files_in_directory('/app/photos/12/'),
            'model': get_files_in_directory('/app/models/12/')[0],
            'badges': ['GreatView', 'SwimmingPool'],
            'description': "Très beau 3 pièces en plein centre de la Ferté Macé, à 10 minutes de Bagnoles de l'Orne. L'appartement est situé au 1 er étage d'une maison de maître, et dispose d'un balcon plein sud, d'une courette privative et d'une place de stationnement. Nombreux éléments d'époque, cheminées, parquets, grande hauteur sous plafond. UN très grand salon et deux belle chambres. L'appartement est très propre, l'ensemble des sols ont été refaits. Idéal personne seule, jeune couple ou retraités. Dossier complet demandé + garant le cés échéant. Premier contact par mail. Disponible début avril.",
            'rooms': 3,
            'bedrooms': 1,
            'furnished': False,
            'energyRating': 'C',
            'rentAmount': 495,
            'chargesAmount': 85
        }
    ]

    properties2 = [
        {
            'area': 83,
            'country': 'France',
            'cityName': 'La Ferté-Macé',
            'address': '60 Avenue Lemeunier de la Raillère',
            'postalCode': '61600',
            'type': 'Appartement',
            'photos': get_files_in_directory('/app/photos/21/'),
            'model': get_files_in_directory('/app/models/21/')[0],
            'badges': ['Terrace', 'Garage', 'Orientation'],
            'description': "Cet appartement, situé à LA FERTE MACE, dispose d'une surface de 80.00m², dont 3 chambre(s). Ce bien possède un parking. Pour un loyer de 429,86€ par mois, ce bien est déjà disponible. Référence annonce : 01302023010015",
            'rooms': 3,
            'bedrooms': 1,
            'furnished': False,
            'energyRating': 'D',
            'rentAmount': 430,
            'chargesAmount': 40
        },
        {
            'area': 83,
            'country': 'France',
            'cityName': 'Javron-les-Chapelles',
            'address': '38 Rue du Dr Cumin',
            'postalCode': '53250',
            'type': 'Maison',
            'photos': get_files_in_directory('/app/photos/22/'),
            'model': get_files_in_directory('/app/models/22/')[0],
            'badges': ['Terrace', 'Garage', 'Orientation'],
            'description': "Maison à louer à Javron-les-Chapelles - Réf 8736 Votre agence Lair Immobilier de Pré-en-Pail vous propose à la location cette maison sur sous-sol complet située à Javron-les-Chapelles au calme mais à proximité des écoles et commerces. Elle se compose au rez-de-chaussée d'une vaste entrée avec placard, wc, séjour salon, cuisine aménagée, deux chambres et une salle de bain. En sous-sol vous trouverez une grande pièce de réception avec cuisine d'été, des rangements, une chambre d'appoint, un garage, une cave et une chaufferie. Jardin d'environ 700 m². Loyer 485euros Dépôt de garantie 485euros Honoraires agence 436.50euros dont 154.20euros pour l'état des lieux d'entrée. Pour plus de renseignements contactez nous au 02.33.28.20.20 A très bientôt! Référence annonce : 8736 Honoraires à la charge du locataire : 436 € TTC dont 282 € pour l'état des lieux Dépôt de garantie : 485 €",
            'rooms': 4,
            'bedrooms': 2,
            'furnished': False,
            'energyRating': 'E',
            'rentAmount': 485,
            'chargesAmount': 120
        },
    ]
    query = gql(
        """
          mutation($input: PropertyInput) {
            createProperty(input: $input) {
              ID
            }
          }
    """
    )

    # Execute the query on the transport
    if count == 1:
        for property in properties1:
            result = connection.client.execute(query, variable_values={
                'input': property
            }, upload_files=True)
            properties.append(result['createProperty'])
    elif count == 2:
        for property in properties2:
            result = connection.client.execute(query, variable_values={
                'input': property
            }, upload_files=True)
            properties.append(result['createProperty'])

    return properties


def create_tenants(count=1):
    tenants = []

    tenants1 = [
        {
            'tel': '0616153745',
            'birthday': '1995-11-12T15:04:05Z',
            'gender': 'MAN',
            'customerRef': '533-11-7524',
            'user': {
                'username': 'john@tek.eu',
                'password': 'john'
            }
        },
        {
            'tel': '206-714-2397',
            'birthday': '1974-10-13T15:04:05Z',
            'gender': 'WOMAN',
            'customerRef': '533-11-3456',
            'user': {
                'username': 'charleen@teleworm.us',
                'password': 'charleen'
            }
        },
    ]

    tenants2 = [
        {
            'tel': '0573043703',
            'birthday': '1939-12-07T15:04:05Z',
            'gender': 'WOMAN',
            'customerRef': '533-11-7396',
            'user': {
                'username': 'ines@rhyta.com',
                'password': 'ines'
            }
        },
        {
            'tel': '0453828191',
            'birthday': '1975-06-30T15:04:05Z',
            'gender': 'MAN',
            'customerRef': '533-11-2073',
            'user': {
                'username': 'channing@dayrep.com',
                'password': 'channing'
            }
        },
    ]
    query = gql(
        """
          mutation($input: TenantInput) {
            createTenantUser(input: $input) {
              ID
              user {
                username
              }
            }
          }
    """
    )

    # Execute the query on the transport
    if count == 1:
        for tenant in tenants1:
            result = connection.client.execute(query, variable_values={
                'input': tenant
            }, upload_files=True)
            result['createTenantUser']['user']['password'] = tenant['user']['password']
            tenants.append(result['createTenantUser'])
    elif count == 2:
        for tenant in tenants2:
            result = connection.client.execute(query, variable_values={
                'input': tenant
            }, upload_files=True)
            result['createTenantUser']['user']['password'] = tenant['user']['password']
            tenants.append(result['createTenantUser'])

    return tenants


def assign_properties(tenants_id, properties_id):
    query = gql(
        """
          mutation assignProperty($tenantId: Int!, $propertyId: Int!) {
            assignProperty(tenantId: $tenantId, propertyId: $propertyId) {
              ID
            }
          }
        """
    )

    index = 0

    while index < 2:
        connection.client.execute(query, variable_values={
            'tenantId': tenants_id[index]['ID'],
            'propertyId': properties_id[index]['ID'],
        }, upload_files=True)
        index += 1


def create_ticket(count):
    tickets = [
        {
            'type': "PAYMENT",
            'description': "Je ne comprend pas pourqoi les charges ont été augmentées ?"
        },
        {
            'type': "MAINTENANCE",
            'description': "⚠️Chauffe eau en panne !⚠️"
        },
        {
            'type': "DOCUMENTS",
            'description': "Demande quittance loyer mois de mai 2020."
        },
        {
            'type': "RENT",
            'description': "Demande de préavis."
        },
    ]
    query = gql(
        """
          mutation($input: AnomalyInput) {
            createAnomaly(input: $input) {
              ID
            }
          }
        """
    )

    result = connection.client.execute(query, variable_values={
        'input': tickets[count]
    }, upload_files=True)
    return result['createAnomaly']


def login_as_estate_agent(estate_agent):
    user = {
        'username': estate_agent['user']['username'],
        'password': estate_agent['user']['password']
    }
    query = gql(
        """
          mutation($input: UserInput) {
            loginAsEstateAgent(input: $input) {
              token
            }
          }
    """
    )

    # Execute the query on the transport
    result = connection.client.execute(query, variable_values={
        'input': user
    })
    return result['loginAsEstateAgent']


def login_as_tenant(tenant):
    user = {
        'username': tenant['user']['username'],
        'password': tenant['user']['password']
    }
    query = gql(
        """
          mutation($input: UserInput) {
            loginAsTenant(input: $input) {
              token
            }
          }
    """
    )

    # Execute the query on the transport
    result = connection.client.execute(query, variable_values={
        'input': user
    })
    return result['loginAsTenant']


def main():
    connection.fetch_connection()

    try:
        company = signup_as_company()
    except TransportQueryError as e:
        if len(e.errors) and e.errors[0]['message'] == 'company seems already register':
            return 0
        else:
            raise e

    print(company)

    connection.fetch_connection(headers={'Authorization': company['token']})

    estate_agents = create_estate_agents()

    print(estate_agents)

    count = 0
    count_tenant = 0
    for estate_agent in estate_agents:
        count += 1

        estate_agent = login_as_estate_agent(estate_agent)

        print(estate_agent)

        connection.fetch_connection(headers={'Authorization': estate_agent['token']})

        properties = create_properties(count)

        print(properties)

        tenants = create_tenants(count)

        print(tenants)

        assign_properties(properties, tenants)

        for tenant in tenants:
            tenant = login_as_tenant(tenant)

            connection.fetch_connection(headers={'Authorization': tenant['token']})

            ticket = create_ticket(count_tenant)

            print(ticket)

            count_tenant += 1


if __name__ == '__main__':
    main()

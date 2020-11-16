
import client from "../apollo_client";
import { gql } from "apollo-boost";
import type { Phone } from "./user";

export interface Label {
  value: string | null;
  locale: string | null;
}

export interface Source {
  url: string | null;
}

export interface Price {
  value: number | null;
  condition: string | null;
  currency: string | null;
}

export interface Asset {
  url: string | null;
}

export interface Geometry {
  type: string | null;
  coordinates: number[] | null;
}

export interface GeoJSON {
  type: string | null;
  geometry: Geometry | null;
}

export interface Location {
  geojson: GeoJSON | null;
}

export interface Surface {
  value: number | null;
  unit: string | null;
  rooms: number | null;
}

export interface Offeror {
  name: string | null;
  email: string | null;
  phone: Phone | null;
  type: string | null;
  prefered_contact_mode: string | null;
}

export interface RentOffer {
  id: string | null;
  title: Label[] | null;
  source: Source | null;
  description: Label[] | null;
  price: Price[] | null;
  assets: Asset[] | null;
  createdAt: string | null;
  offeror: Offeror | null;
  location: Location | null;
  surface: Surface | null;
}

export interface SearchQuery {
  limit: number | null;
  offset: number | null;
  count: number | null;
  intext: string | null;
}

const RENT_OFFER_QUERY = gql`
query {
  rentoffers {
    id
    surface {
      value
      unit
    }
    price {
      value
      currency
    }
    title {
      value
    }
    assets {
      url
    }
    description {
      value
    }
    source {
      url
    }
    offeror {
      id
      name
      email
      phone {
        countryCode
        value
      }
    }
  }
}
`

export const GET_RENTOFFERS = async (query?: SearchQuery): Promise<RentOffer[]> => {
  let rt = await client.query<RentOffer[]>({
    query: RENT_OFFER_QUERY,
    variables: query,
  })
  return rt.data;
}
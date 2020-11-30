export interface ApiEnvironnement {
  name: string;
  hostUrl: string;
}

export interface Environnement {
  apiEnvironnements: ApiEnvironnement[];
  canChangeEnvironnement: boolean;

  loadSavedUrl(): string;
}

abstract class BaseEnvironnement implements Environnement {
  public abstract apiEnvironnements: ApiEnvironnement[];
  public abstract canChangeEnvironnement: boolean;

  public loadSavedUrl(): string {
      const savedEnvironnement = localStorage.getItem("environnement");
      return savedEnvironnement ? ((JSON.parse(savedEnvironnement) as ApiEnvironnement).hostUrl) : `${process.env.VUE_APP_API_HOST_URL}`;
  }
}


class LocalEnvironnement extends BaseEnvironnement {
  public canChangeEnvironnement: boolean = true;
  public apiEnvironnements: ApiEnvironnement[] = [
      { name: 'Local', hostUrl: "http://localhost:3001" },
  ];
}

class ProdEnvironnement extends BaseEnvironnement {
  public canChangeEnvironnement: boolean = false;
  public apiEnvironnements: ApiEnvironnement[] = [
      { name: process.env.VUE_APP_ENV_NAME, hostUrl: process.env.VUE_APP_API_HOST_URL }
  ];
}

const environnement: Environnement = process.env.NODE_ENV === 'production' ? new ProdEnvironnement() : new LocalEnvironnement();
export default environnement;

import axios, { AxiosInstance, AxiosResponse, AxiosRequestConfig } from 'axios';
import { AuthentificationRequest } from './types';

const hostUrl: string = `${process.env.VUE_APP_API_HOST_URL}`;
const tokenName: string = "token";

class FetchV2Service {
    private http!: AxiosInstance;
    private token!: string;

    constructor() {
        this.http = axios.create({
            baseURL: `${hostUrl}`,
        });
    }

    get requestConfig() {
        return {
            headers: {
                Authorization: `${window.localStorage.getItem(tokenName)}`
            }
        };
    }

    public get(url: string, params: any = {}) {
        return new Promise<any>((resolve, reject) => {
            this.http.get(url, {
                params,
                ...this.requestConfig
            }).then(resp => {
                resolve(resp.data);
            }).catch(e => {
                this.handleError(e, reject);
            });
        });
    }

    public post(url: string, data: any) {
        return new Promise((resolve, reject) => {
            this.http.post(url, data, this.requestConfig).then(resp => {
                resolve(resp.data);
            }).catch(e => {
                this.handleError(e, reject);
            });
        });
    }

    public authenticate(authConfig: AuthentificationRequest): Promise<any> {
        return new Promise((resolve, reject) => {
            this.http.post('login', authConfig).then(resp => {
                const token = resp.data.data.token;
                localStorage.setItem(tokenName, token);
                resolve({ user: resp.data.data, token });
            }).catch(reject);
        });
    }

    public put(url: string, data: any, additionalHeaders: { [key: string]: any } = {}, responseMapper?: (resp: AxiosResponse<any>) => any): Promise<any> {
        return new Promise((resolve, reject) => {
            const config: AxiosRequestConfig = {
                ...this.requestConfig,
                headers: {
                    ...this.requestConfig.headers,
                    ...additionalHeaders
                }
            };

            this.http.put(url, data, config).then(resp => {
                this.saveToken(resp);
                const result = (responseMapper && responseMapper(resp)) || resp.data;
                resolve(result);
            }).catch(e => {
                this.handleError(e, reject);
            });
        });
    }

    public delete(url: string) {
        return new Promise((resolve, reject) => {
            this.http.delete(url, this.requestConfig).then(resp => {
                this.saveToken(resp);
                resolve(resp.data);
            }).catch(e => {
                this.handleError(e, reject);
            });
        });
    }

    private saveToken(response: AxiosResponse<any>) {
        this.token = response.data.token;
        if (this.token) {
            window.localStorage.setItem(tokenName, this.token);
        }
    }

    private handleError(e: any, reject: any) {
        reject(e);
    }
}

export default new FetchV2Service();

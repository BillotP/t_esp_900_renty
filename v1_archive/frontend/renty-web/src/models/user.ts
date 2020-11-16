export interface Email {
    value: string | null;
    valid: boolean | null;
}

export interface Phone {
    value: string | null;
    valid: boolean | null;
    countryCode: string | null;
}

export interface User {
    id: string | null;
    pseudo: string | null;
    gender: string | null;
    firstName: string | null;
    lastName: string | null;
    email: Email | null;
    phone: Phone | null;
    password: string | null;
}
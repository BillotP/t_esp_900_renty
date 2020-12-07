export type Rule = (v: any) => boolean | string;
export const RequiredRule: Rule = (v: any) => !!v || "Ce champ est obligatoire";
export const MailFormatRule: Rule = (v: any) =>
    /.+@.+/.test(v) || "L'email est incorrect";
export const PasswordRule: Rule = (v: any) =>
    (v && /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*(),.?":{}|<>])[0-9a-zA-Z!@#$%^&*(),.?":{}|<>]{8,}$/.test(v))
    || "Le mot de passe doit avoir au moins 8 caratères dont une majuscule, un chiffre et un caractère spécial";

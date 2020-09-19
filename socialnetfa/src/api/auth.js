import { API_HOST, TOKEN } from "../utils/constants";

export function signUpApi(user) {
    const url = `${API_HOST}/registro`;
    const userTemp = {
        ...user,
        email: user.email.toLowerCase(),
        fechaNacimiento: new Date()
    };
    delete userTemp.repeatPassword;

    const params = {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(userTemp)
    };

    return fetch(url, params).then(response => {
        if(response.status >= 200 && response.status < 300){
            return response.json();
        }
        return { code: 404, message: "Email no disponible" }
    }).then(result => {
        return result;
    }).catch(err => {
        return err;
    })
}

export function signInApi(user) {
    const url = `${API_HOST}/login`;
    const userTemp = {
        ...user,
        email: user.email.toLowerCase()
    };

    const params = {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(userTemp)
    };

    return fetch(url, params).then(response => {
        if(response.status >= 200 && response.status < 300){
            return response.json();
        }
        return { code: 404, message: "Email o Password incorrectos" }
    }).then(result => {
        return result;
    }).catch(err => {
        return err;
    })
}

export function setTokenApi(token){
    localStorage.setItem(TOKEN, token);
}
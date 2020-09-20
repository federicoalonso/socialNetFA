import { API_HOST } from '../utils/constants';
import { getTokenApi } from './auth';

export function getUserApi(id) {
    const url = `${API_HOST}/perfil?id=${id}`;

    const params = {
        headers: {
            "Content-Type": "application/json",
            Authorization: `${getTokenApi()}`
        }
    }

    return fetch(url, params)
        .then(response => {
            if (response.status >= 400) throw null;
            return response.json();
        })
        .then(result => {
            return result;
        })
        .catch(err => {
            return err;
        });
}
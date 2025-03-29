import axios from 'axios'

export const postLogin = (username: string, password: string) =>
    axios.post("http://192.168.2.193:3000/login", { username, password })
export const postRegister = (username: string, password: string) =>
    axios.post("http://192.168.2.193:3000/register", { username, password })

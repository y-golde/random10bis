import axios from 'axios';

const setUpAxios = () => {
    axios.defaults.baseURL = '/api/'; //TODO: env
};

export default setUpAxios;

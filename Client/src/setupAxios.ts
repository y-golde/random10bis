import axios from 'axios';

const setUpAxios = () => {
    axios.defaults.baseURL = 'http://localhost:9000/'; // '/api/'; //TODO: env
};

export default setUpAxios;

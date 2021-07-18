import App from './App.svelte';
import setUpAxios from './setupAxios';

setUpAxios();

var app = new App({
    target: document.body,
});

export default app;

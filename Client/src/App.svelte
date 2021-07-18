<script lang="ts">
    import axios from 'axios';
    import { onMount } from 'svelte';

    import { Menu } from './Types/Menu';
    import MenuComponent from './Components/Menu/Menu.svelte';

    let menu: Menu;

    const fetchMenu = async (restaurantId: string) => {
        const res = await axios.get<Menu>('/menu', { params: { restaurantId } });
        if (res.status === 200) {
            menu = res.data;
        }

        return menu;
    };

    onMount(async () => {
        menu = await fetchMenu('29321');
    });
</script>

<h1>תן בי(ס)</h1>
{console.log(menu)}
{#if menu}
    <MenuComponent menu="{menu}" />
{/if}

<style>
    h1 {
        text-align: center;
    }
</style>

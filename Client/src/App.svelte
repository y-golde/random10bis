<script lang="ts">
    import axios from 'axios';
    import { onMount } from 'svelte';

    import type Menu from './Types/Menu';
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

<div class="text-4xl text-center">תן בי(ס)</div>

{#if menu}
    <MenuComponent menu="{menu}" />
{/if}

<style global lang="postcss">
    @tailwind base;
    @tailwind components;
    @tailwind utilities;
</style>

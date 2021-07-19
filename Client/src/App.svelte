<script lang="ts">
    import { onMount } from 'svelte';

    import fetchMenu from './fetchMenu';
    import type Menu from './Types/Menu';
    import MenuComponent from './Components/Menu/Menu.svelte';

    let menu: Menu | null = null;
    let selectedId = 3;

    onMount(async () => {
        menu = await fetchMenu('29321');
    });

    const getRandomValidId = () => {
        const allValidIds = menu.categoriesList.flatMap((category) => category.dishList.map((dish) => dish.id));

        selectedId = allValidIds[Math.floor(Math.random() * allValidIds.length)];
    };
</script>

<div class="text-4xl text-center">תן בי(ס)</div>

{#if menu}
    <MenuComponent menu="{menu}" selectedId="{selectedId}" />
    <button on:click="{getRandomValidId}">clickme</button>
{/if}

<style global lang="postcss">
    @tailwind base;
    @tailwind components;
    @tailwind utilities;
</style>

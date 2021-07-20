<script lang="ts">
    import fetchMenu from './fetchMenu';
    import type Menu from './Types/Menu';
    import MenuComponent from './Components/Menu/Menu.svelte';

    let menu: Menu | null = null;
    let selectedId = 3;
    let menuId = '';

    const setRandomValidId = () => {
        const allValidIds = menu.categoriesList.flatMap((category) => category.dishList.map((dish) => dish.id));

        selectedId = allValidIds[Math.floor(Math.random() * allValidIds.length)];
    };

    const setMenu = async () => {
        menu = null;
        if (menuId) {
            menu = await fetchMenu(menuId);
            setRandomValidId();
        }
    };
</script>

<div class="text-4xl text-center">תן בי(ס)</div>

<div>
    <input
        class="shadow appearance-none border rounded py-2 px-3 text-gray-700 focus:outline-none focus:shadow-outline"
        type="text"
        placeholder="restaurantId"
        bind:value="{menuId}"
    />
    <button class="btn bg-blue-500 rounded-full p-2 text-white" on:click="{setMenu}">import resteraunt menu</button>
</div>

{#if menu}
    <button class="btn bg-blue-500 rounded-full p-2" on:click="{setRandomValidId}">Choose a random food!</button>
    <MenuComponent menu="{menu}" selectedId="{selectedId}" />
{/if}

<style global lang="postcss">
    @tailwind base;
    @tailwind components;
    @tailwind utilities;
</style>

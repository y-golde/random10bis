<script lang="ts">
    import type Menu from '../../Types/Menu';
    import Category from './Category/Category.svelte';

    export let menu: Menu;
    export let selectedId: number;

    let categoriesList: Category[] = [];

    $: {
        categoriesList = menu.categoriesList as any as Category[];
    }
</script>

<div class="menu container mx-auto">
    {#each categoriesList as cat}
        <Category category="{cat}" selectedId="{selectedId}" />
    {/each}
</div>

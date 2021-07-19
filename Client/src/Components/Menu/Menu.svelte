<script lang="ts">
    import type Menu from '../../Types/Menu';
    import Category from './Category/Category.svelte';

    export let menu: Menu;
    export let selectedId: number;

    const { categoriesList } = menu;
</script>

<div class="menu container mx-auto">
    {#each categoriesList as category}
        <Category category="{category}" selectedId="{selectedId}" />
    {/each}
</div>

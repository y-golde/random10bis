<script lang="ts">
    import type Menu from '../../Types/Menu';
    import Category from './Category/Category.svelte';

    export let menu: Menu;

    const { categoriesList } = menu;
</script>

<div class="menu">
    {#each categoriesList as category}
        <Category category="{category}" />
    {/each}
</div>

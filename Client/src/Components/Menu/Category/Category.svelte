<script lang="ts">
    import Dish from './Dish/Dish.svelte';
    import type Category from '../../../Types/Category';

    export let selectedId: number;
    export let category: Category;
</script>

<div class="text-right text-2xl">
    {category.categoryName}
    {#if category.categoryDesc != ''}
        <span class="text-base mr-3">{category.categoryDesc}</span>
    {/if}
</div>
<div class="container mx-auto rtl grid grid-cols-3 gap-4 flex-col-reverse my-2">
    {#each category.dishList as dish}
        <Dish dish="{dish}" selectedId="{selectedId}" />
    {/each}
</div>

<style>
    .rtl {
        direction: rtl;
    }
</style>

<script lang="ts">
    import type Category from '../../../Types/Category';
    export let category: Category;
</script>

<h2>{category.categoryName}</h2>
<h2>{category.categoryDesc}</h2>

<style>
    h2 {
        text-align: center;
    }
</style>

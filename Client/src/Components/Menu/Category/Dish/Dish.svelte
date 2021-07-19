<script lang="ts">
    import type Dish from '../../../../Types/Dish';

    export let dish: Dish;
    export let selectedId: number;

    const { dishName, dishDescription, dishPrice, id } = dish;

    let selectedClass = '';
    $: {
        const isSelected = id === selectedId;
        selectedClass = isSelected ? 'font-bold' : '';
    }
</script>

<div class="col-span-1 text-xl {selectedClass}">
    {dishName}
    {dishDescription}
    {dishPrice}
</div>

<style>
    .selected {
        border-radius: 15px;
        border: 3px solid gold;
    }
</style>

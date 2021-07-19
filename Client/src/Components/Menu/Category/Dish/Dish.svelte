<script lang="ts">
    import type Dish from '../../../../Types/Dish';

    export let dish: Dish;

    const { dishName, dishDescription, dishPrice } = dish;
</script>

<div class="col-span-1 text-xl">
    {dishName}
    {dishDescription}
    {dishPrice}
</div>

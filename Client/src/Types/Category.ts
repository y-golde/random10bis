import type Dish from './Dish';

export interface Category {
    categoryName: string;
    categoryDesc: string;
    dishList: Dish[];
}

export default Category;

import type Dish from './Dish';

export interface Category {
    valid: boolean;
    categoryName: string;
    categoryDesc: string;
    dishList: Dish[];
}

export default Category;

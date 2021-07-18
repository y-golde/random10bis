export interface Menu {
    categoriesList: Category[];
}

export interface Category {
    categoryName: string;
    categoryDesc: string;
    dishList: Dish[];
}

export interface Dish {
    dishName: string;
    isPopularDish: boolean;
    dishPopularityScore: number;
    dishPrice: number;
    dishDescription: string;
}

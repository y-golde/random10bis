import axios from 'axios';

import type Menu from './Types/Menu';

const fetchMenu = async (restaurantId: string) => {
    const res = await axios.get<Menu>('/menu', { params: { restaurantId } });
    if (res.status === 200) {
        return addIdsToMenu(res.data);
    }

    return null;
};

const addIdsToMenu = (menu: Menu) => {
    let i = 1;
    menu.categoriesList.forEach((category) =>
        category.dishList.forEach((dish) => {
            dish.id = i++;
        })
    );

    return menu;
};

export default fetchMenu;

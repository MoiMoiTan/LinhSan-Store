// API liên quan đến sản phẩm
import axios from 'axios';

const getAllProducts = async () => {
    const response = await axios.get('/api/products');
    return response.data;
};

export { getAllProducts };

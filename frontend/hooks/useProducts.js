// Hook quản lý sản phẩm
import { useState, useEffect } from 'react';
import { getAllProducts } from '../services/productService';

const useProducts = () => {
    const [products, setProducts] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        getAllProducts().then(data => {
            setProducts(data);
            setLoading(false);
        });
    }, []);

    return { products, loading };
};

export default useProducts;

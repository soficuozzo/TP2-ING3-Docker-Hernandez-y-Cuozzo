import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import toast, { Toaster } from 'react-hot-toast';

const CerrarSesion = () => {
    const navigate = useNavigate();

    useEffect(() => {

        const loadingToastId = toast.success('Cerrando sesión...');

        // Limpiar el almacenamiento local
        localStorage.removeItem('token');
        localStorage.removeItem('usuario_id');
        localStorage.removeItem('email');
        localStorage.removeItem('nombre');
        localStorage.removeItem('tipo');

      
        setTimeout(() => {
            toast.dismiss(loadingToastId); // Cerrar el toast de "Cerrando sesión..."
            navigate('/'); // Redirigir a la página de inicio
        }, 2000);

    }, [navigate, ]);

    return (
        <Toaster />
    );
};

export default CerrarSesion;

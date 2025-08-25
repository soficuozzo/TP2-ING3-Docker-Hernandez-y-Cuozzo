import React, { useState, useEffect } from 'react';
import './resultados.css';
import toast, { Toaster } from 'react-hot-toast';
import axios from 'axios';
import 'react-responsive-carousel/lib/styles/carousel.min.css';
import './cursos.css';
import { useNavigate } from "react-router-dom";

const Contenedores = () => {
    const navigate = useNavigate();   
     const [contenedores, setContenedores] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);


    useEffect(() => {
        const fetchContenedores = async () => {

            setLoading(true);
            
            try {
                const response = await axios.get('http://localhost:8086/containers');
                console.log('Respuesta de la API:', response); 
                setContenedores(response.data);


            } catch (error) {
                setError('Error fetching contenedores');
                console.error('Error fetching contenedores:', error);
            } finally {
                setLoading(false);
            }
        };

        fetchContenedores();
    }, []);

    if (loading) {
        return <div>Cargando contenedores...</div>;
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

   const start = (status) => {
        if(status == "running") return false;
        return true;
    };


    const handleStart = async (name) => {
      
        try {
            const response = await axios.post(`http://localhost:8086/containers/start/${name}`, {});
            if (response.status === 200) {
            toast.success('Api corriendo');
                

            } else {
                alert('Error al correr api');
            }
        } catch (error) {
            console.error('Error al correr api:', error);
            if (error.response) {
                console.error('Detalles del error:', error.response.data);
            }
            alert('Error al correr api. Verifica tu conexión o inténtalo más tarde.');
        }

        navigate(0);

    };

    const handleExit = async (name) => {
      
        try {
            const response = await axios.post(`http://localhost:8086/containers/stop/${name}`, {});
            if (response.status === 200) {    
                toast.success('API en stop');
            } else {
                alert('Error al parar api');
            }
        } catch (error) {
            console.error('Error al parar api:', error);
            if (error.response) {
                console.error('Detalles del error:', error.response.data);
            }
            alert('Error al parar la api. Verifica tu conexión o inténtalo más tarde.');
        }

        navigate(0);          

    };



    return (
        <div className="fullscreen-container">
            
    
            <div className="Resultadosfondo">
                
                
            
                <div className="course-list">
                    {contenedores.length > 0 ? (
                        contenedores.map(contenedor => (
                            <div key={contenedor.name} className="course-list-item">
                                <div>
                                    <strong>{contenedor.name}</strong>
                                    <p>{contenedor.status}</p>
                                    <br/>

                            {start(contenedor.status) && (
                                <>
                                    <button onClick={() => handleStart(contenedor.name)} className="inscribirsebutton">Start</button>
                                </>
                            )}

                            {!start(contenedor.status) && (
                                <>
                                    <button onClick={() => handleExit(contenedor.name)} className="inscribirsebutton">Exit</button>
                                </>
                            )}  
                            



                                </div>
                                <hr />
                            </div>
                        ))
                    ) : (
                        <p>No se encontraron contenedores.</p>
                    )}
                </div>

            </div>
        </div>
    
    );
};

export default Contenedores;

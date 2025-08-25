import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './resultados.css';

import axios from 'axios';
import 'react-responsive-carousel/lib/styles/carousel.min.css';
import './cursos.css';
import { useParams } from 'react-router-dom';

const Cursos = () => {
    const { curso_id } = useParams();
    const [cursos, setCursos] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const usuariotipo = localStorage.getItem("tipo");


    useEffect(() => {
        const fetchCursos = async () => {

            setLoading(true);
            
            try {
                const response = await axios.get('http://localhost:8082/cursos/todos');
                console.log('Respuesta de la API:', response); 
                setCursos(response.data);


            } catch (error) {
                setError('Error fetching cursos');
                console.error('Error fetching cursos:', error);
            } finally {
                setLoading(false);
            }
        };

        fetchCursos();
    }, []);

    if (loading) {
        return <div>Cargando cursos...</div>;
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

   

    return (
        <div className="fullscreen-container">
            
    
            <div className="Resultadosfondo">
                
                
            
                <div className="course-list">
                    {cursos.length > 0 ? (
                        cursos.map(curso => (
                            <div key={curso.id} className="course-list-item">
                                <div>
                                    <Link to={`/curso/${curso.id}`} className="curso-link">
                                        <strong>{curso.nombre}</strong>
                                    </Link>
                                    <p>{curso.descripcion}</p>
                                    <p>{curso.categoria}</p>
                                    <br/>
                                    Para saber más, haga click al nombre del curso.
                                </div>
                                <hr />
                            </div>
                        ))
                    ) : (
                        <p>No se encontraron cursos.</p>
                    )}
                </div>

            </div>
        </div>
    
    );
};

export default Cursos;

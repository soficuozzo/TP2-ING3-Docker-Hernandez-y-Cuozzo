import React, { useState, useEffect } from 'react';
import axios from 'axios';
import toast, { Toaster } from 'react-hot-toast';
import { useParams } from 'react-router-dom';
import './resultados.css';
import './curso.css';
import { useNavigate } from "react-router-dom";

const Curso = () => {
    const { curso_id } = useParams();
    console.log("Curso ID recibido:", curso_id);

    const [curso, setCurso] = useState(null);
    const [loading, setLoading] = useState(true);
    const [estaInscripto, setEstaInscripto] = useState(false);
    const [dispo, setDispo] = useState(false);
    const usuariotipo = localStorage.getItem("tipo");
    const [error, setError] = useState(null);

    const navigate = useNavigate(); // Permite la navegación entre páginas con las rutas

    useEffect(() => {
        const fetchCurso = async () => {
            try {
                const response = await axios.get(`http://localhost:8082/cursos/${curso_id}`);
                setCurso(response.data);
            } catch (error) {
                setError('Error fetching curso');
                console.error('Error fetching curso:', error);
            } finally {
                setLoading(false);
            }
        };

        fetchCurso();
    }, [curso_id]);

   

    useEffect(() => {
        const verificarInscripcion = async () => {
            const usuarioId = localStorage.getItem('usuario_id');
            if (!usuarioId) {
                return;
            }

            try {
                const response = await axios.get(`http://localhost:8082/usuario/miscursos/${usuarioId}`);
                const cursosInscritos = response.data.cursos;
                const estaInscritoEnCurso = cursosInscritos.some(cursoInscrito => cursoInscrito.id === curso_id);
                setEstaInscripto(estaInscritoEnCurso);

            } catch (error) {
                console.error('Error al obtener cursos inscritos:', error);
            }
        };

        verificarInscripcion();
    
    }, [curso_id]);

    const handleInscripcion = async () => {
        const usuarioId = localStorage.getItem('usuario_id');
        if (!usuarioId) {
            alert('No se encontró el ID de usuario en el almacenamiento local.');
            return;
        }

        try {
            const response = await axios.post(`http://localhost:8082/inscripcion`, {
                usuario_id: parseInt(usuarioId, 10),
                curso_id: curso_id
            });
            if (response.status === 201) {
                toast.success('Inscripción exitosa');
                setEstaInscripto(true);
            } else {
                alert('Error al inscribirse');
            }
        } catch (error) {
            console.error('Error al inscribirse:', error);
            if (error.response) {
                console.error('Detalles del error:', error.response.data);
            }
            alert('Error al inscribirse. Verifica tu conexión o inténtalo más tarde.');
        }
    };


    const mostrar1 = () => {
        if (!localStorage.getItem("usuario_id")) { return false; }
        if (usuariotipo === "admin") { return false; }
        if (estaInscripto) { return false; }
        if(curso.capacidad === 0){return false}
        return true;
    };

    const mostrar2 = () => {
        if (usuariotipo === "estudiante") { return false; }
        return true;
    };
    

    if (loading) {  
        return <div>Cargando curso...</div>;
    }


    if (!curso) {
        return <div>No se encontró el curso seleccionado.</div>;
    }

    const handleUpdate = () => {
        navigate(`/updatecurso/${curso_id}`); 
        localStorage.setItem("cursoid", curso.id)

      };

   

    return (
        
        <div className="fullscreen-container">
            <div className="Resultadosfondo">
                <div className="course-list-item">
                    <h3>{curso.nombre}</h3>
                    <p>{curso.descripcion}</p>
                    <p><b>Categoria:</b> {curso.categoria}</p>
                    <p><b>Profesor:</b> {curso.profesor}</p>
                    <p><b>Duracion:</b> {curso.duracion}</p>
                    <p><b>Requisitos:</b> {curso.requisito}</p>
                    <p><b>Capacidad:</b> {curso.capacidad}</p> 
                    
                    {mostrar1() && (
                        <>
                            <button onClick={handleInscripcion} className="inscribirsebutton">Inscribirme</button>
                            <Toaster position="bottom-center" />
                        </>
                    )}
                    <br/><br/>

                    {mostrar2() && (
                        <>
                            <button onClick={handleUpdate} className="inscribirsebutton">Modificar</button>
                            <Toaster position="bottom-center" />
                        </>
                    )}
                    
                    <br/>

                </div>
                <br/><br/><br/>

                

                <Toaster position="bottom-center" />
            </div>
        </div>
    );
};

export default Curso;

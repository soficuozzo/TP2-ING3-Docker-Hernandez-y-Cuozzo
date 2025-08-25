import React, { useState } from 'react';
import { useNavigate } from "react-router-dom";
import '../componentes/NavegadorHome.css';
import toast, { Toaster } from 'react-hot-toast';
import './registrarse.css'


const CursoNuevo = () => {
    const [nombre, setNombre] = useState("");
    const [descripcion, setDescripcion] = useState("");
    const [categoria, setCategoria] = useState("");
    const [capacidad, setCapacidad] = useState("");
    const [profesor, setProfesor] = useState("");
    const [duracion, setDuracion] = useState("");
    const [requisitos, setRequisitos] = useState("");
    const [valoracion, setValoracion] = useState("");

    const imagen = "foto";

    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        
        let valid = true

        if (nombre === '') {
            document.getElementById("inputNombre").style.borderColor = 'red';
            toast.error("Ingrese el nombre correctamente."); // Muestra un mensaje de error si las credenciales son incorrectas
            valid = false
        }

        if (descripcion === ''){
            document.getElementById("inputDescripcion").style.borderColor = 'red';
            toast.error("Ingrese la descripcion correctamente."); // Muestra un mensaje de error si las credenciales son incorrectas
            valid = false

        }

        if (categoria === ''){
            document.getElementById("inputCategoria").style.borderColor = 'red';
            toast.error("Ingrese la categoria."); // Muestra un mensaje de error si las credenciales son incorrectas
            valid = false

        }

        if (capacidad === ''){
            document.getElementById("inputCapacidad").style.borderColor = 'red';
            toast.error("Ingrese la capacidad."); // Muestra un mensaje de error si las credenciales son incorrectas
            valid = false

        }

        if (profesor === ''){
            document.getElementById("inputProfesor").style.borderColor = 'red';
            toast.error("Ingrese el profesor."); // Muestra un mensaje de error si las credenciales son incorrectas
            valid = false

        }

        if (duracion === ''){
            document.getElementById("inputDuracion").style.borderColor = 'red';
            toast.error("Ingrese la duracion de cursada."); // Muestra un mensaje de error si las credenciales son incorrectas
            valid = false
        }

        if (requisitos === ''){
            document.getElementById("inputRequisitos").style.borderColor = 'red';
            toast.error("Ingrese los requisitos."); // Muestra un mensaje de error si las credenciales son incorrectas
            valid = false
        }

        if (valoracion === ''){
            document.getElementById("inputValoracion").style.borderColor = 'red';
            toast.error("Ingrese la valoracion."); // Muestra un mensaje de error si las credenciales son incorrectas
            valid = false
        }


        if(valid){
                try{
                
                    const check = await fetch(`http://localhost:8082/cursos/nombre/${nombre}`)
                    if(check.ok){
                        toast.error("Ya existe un curso con ese nombre.");
                    }else{
                        const response = await fetch('http://localhost:8082/cursos/create', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
    
                            },
                            body: JSON.stringify({ 
                                nombre, 
                                descripcion, 
                                categoria,
                                capacidad: parseInt(capacidad), 
                                imagen, 
                                profesor, 
                                duracion: parseInt(duracion),
                                requisitos, 
                                valoracion: parseInt(valoracion) }),
    
                        });
    
                        if (response.ok) {
                    
                            navigate('/cursos/todos'); // Redirige al usuario a la página principal después del login exitoso
                            toast.success("Creación exitosa.");
    
                        } else {
                            toast.error(    "Curso Invalido");
                        }
                    }

                }catch (error){
                    console.error('Error al realizar la solicitud al backend:', error);
                    toast.error("Error al intentar crear el curso ");
                }
        }   
    };


    return (
        <div className="login-form-container">
            <form onSubmit={handleSubmit} className="login-form">

                <div className="form-group">
                    <label>Nombre</label> 
                    <br/>
                    <input
                        id={"inputNombre"}
                        type="text"
                        value={nombre}
                        onChange={(e) => setNombre(e.target.value)}
                        placeholder="Ingrese el nombre"
                        
                    />
                </div>

                <div className="form-group">
                    <label>Descripcion</label> 
                    <br/>
                    <input
                        id={"inputDescripcion"}
                        type="text"
                        value={descripcion}
                        onChange={(e) => setDescripcion(e.target.value)}
                        placeholder="Ingrese el nombre"
                        
                    />
                </div>

                <div className="form-group">
                <br/>
                    <label>Categoria</label>
                    <br/>
                    <input
                        id={"inputCategoria"}
                        type="text"
                        value={categoria}
                        onChange={(e) => setCategoria(e.target.value)}
                        placeholder="Ingrese la categoria"
                        
                    />
                </div>
                <div className="form-group">
                <br/>
                    <label>Capacidad</label>
                    <br/>
                    <input
                        id={"inputCapacidad"}
                        type="text"
                        value={capacidad}
                        onChange={(e) => setCapacidad(e.target.value)}
                        placeholder="Ingrese la capacidad"
                        
                    />
                </div>

                <div className="form-group">
                <br/>
                    <label>Profesor</label>
                    <br/>
                    <input
                        id={"inputProfesor"}
                        type="text"
                        value={profesor}
                        onChange={(e) => setProfesor(e.target.value)}
                        placeholder="Ingrese el profesor"
                        
                    />
                </div>

                <div className="form-group">
                <br/>
                    <label>Duracion</label>
                    <br/>
                    <input
                        id={"inputDuracion"}
                        type="text"
                        value={duracion}
                        onChange={(e) => setDuracion(e.target.value)}
                        placeholder="Ingrese la duración de cursada"
                        
                    />
                </div>

                <div className="form-group">
                <br/>
                    <label>Requisitos</label>
                    <br/>
                    <input
                        id={"inputRequisitos"}
                        type="text"
                        value={requisitos}
                        onChange={(e) => setRequisitos(e.target.value)}
                        placeholder="Ingrese los requisitos"
                    />
                </div>

                <div className="form-group">
                <br/>
                    <label>Valoracion</label>
                    <br/>
                    <input
                        id={"inputValoracion"}
                        type="text"
                        value={valoracion}
                        onChange={(e) => setValoracion(e.target.value)}
                        placeholder="Ingrese la valoracion"
                    />
                </div>

                <br/><br/>

                <button type="submit">Crear Curso</button>
            </form>
            <Toaster /> {/* Componente para mostrar notificaciones toast */}
        </div>
    );
};
export default CursoNuevo;

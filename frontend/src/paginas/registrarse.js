import React, { useState } from 'react';
import { useNavigate } from "react-router-dom";
import '../componentes/NavegadorHome.css';
import toast, { Toaster } from 'react-hot-toast';
import './registrarse.css'

const fetchUserData = async (email) => {
    try {
        const response = await fetch(`http://localhost:8081/users/email/${email}`);
        if (response.ok) {
            return await response.json(); // Devuelve los datos del usuario si la solicitud es exitosa
        } else {
            console.error('Error al obtener datos del usuario:', response.status);
            return null; // Devuelve null si la solicitud no es exitosa
        }
    } catch (error) {
        console.error('Error al realizar la solicitud al backend:', error);
        return null; // Devuelve null si hay un error durante la solicitud
    }
};



const Registrarse = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [nombre, setNombre] = useState("");
    const [apellido, setApellido] = useState("");
    const [tipo, setTipo] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        
        let valid = true

        if (email === '') {
            document.getElementById("inputEmailReg").style.borderColor = 'red';
            toast.error("Ingrese el email correctamente."); // Muestra un mensaje de error si las credenciales son incorrectas
            valid = false
        }

        if (password === ''){
            document.getElementById("inputPasswordReg").style.borderColor = 'red';
            toast.error("Ingrese la contraseña correctamente."); // Muestra un mensaje de error si las credenciales son incorrectas
            valid = false

        }

        if (nombre === ''){
            document.getElementById("inputNombreReg").style.borderColor = 'red';
            toast.error("Ingrese su nombre."); // Muestra un mensaje de error si las credenciales son incorrectas
            valid = false

        }

        if (apellido === ''){
            document.getElementById("inputApellidoReg").style.borderColor = 'red';
            toast.error("Ingrese su apellido."); // Muestra un mensaje de error si las credenciales son incorrectas
            valid = false

        }

        if(valid){
                try{
                
                    const check = await fetch(`http://localhost:8081/users/email/${email}`)

                    if(check.ok){
                        toast.error("Ya existe un usuario con ese email.");
                    }else{
                        const response = await fetch('http://localhost:8081/users/create', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
    
                            },
                            body: JSON.stringify({ email, password, tipo, nombre, apellido }),
    
                        });
    
                        if (response.ok) {
                    
                            navigate('/login'); // Redirige al usuario a la página principal después del login exitoso
                            toast.success("Creación exitosa. Ingrese nuevamente.");
    
                        } else {
                            toast.error(    "Usuario Invalido");
                        }
                    }

                }catch (error){
                    console.error('Error al realizar la solicitud al backend:', error);
                    toast.error("Error al intentar iniciar sesión");
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
                        id={"inputNombreReg"}
                        type="text"
                        value={nombre}
                        onChange={(e) => setNombre(e.target.value)}
                        placeholder="Ingrese su nombre"
                        
                    />
                </div>

                <div className="form-group">
                <br/>
                    <label>Apellido</label>
                    <br/>
                    <input
                        id={"inputApellidoReg"}
                        type="text"
                        value={apellido}
                        onChange={(e) => setApellido(e.target.value)}
                        placeholder="Ingrese su apellido"
                        
                    />
                </div>

                <div className="form-group">
                <br/>
                    <label>Email</label>
                    <br/>
                    <input
                        id={"inputEmailReg"}
                        type="text"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        placeholder="Ingrese su email"
                        
                    />
                </div>

                <div className="form-group">
                <br/>
                    <label>Tipo</label>
                    <br/>
                    <select value={tipo}
                        onChange={(e) => setTipo(e.target.value)}
                        
                        id={"inputTipoReg"}
                        >                       
                         <option value="">Elegir</option>
                        <option value="estudiante">Estudiante</option>
                        <option value="admin">Admin</option>
                        

                    </select>
                    
                </div>

                <div className="form-group">
                <br/>
                    <label>Contraseña</label>
                    <br/>
                    <input
                        id={"inputPasswordReg"}
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        placeholder="Ingrese su contraseña"
                    />
                </div>

                <br/><br/>

                <button type="submit">Registrarse</button>
            </form>
            <Toaster /> {/* Componente para mostrar notificaciones toast */}
        </div>
    );
};

export default Registrarse;

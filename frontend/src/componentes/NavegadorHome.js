import React, { useState } from 'react';
import { Link, useNavigate  } from 'react-router-dom';
import './NavegadorHome.css';

const NavegadorHome = ({ onSearch }) => {
  const [palabra, setSearchTerm] = useState('');
  const navigate = useNavigate();

  const handleSearchChange = (event) => {
    setSearchTerm(event.target.value);
  };

  const handleSearchSubmit = (event) => {
    event.preventDefault();
    onSearch(palabra);
    navigate(`/cursos/buscar?palabra=${palabra}`);
    setSearchTerm(''); 
    // Limpiar el término de búsqueda después de buscar
  };

  const usuariotipo = localStorage.getItem("tipo");
  


  return (
    <nav className="navbar">
      <h2>CodeWave Learning</h2>
      <ul className="navlinks">

        <li><Link to="/">Home</Link></li>
        
        <li><Link to="/cursos/todos">Cursos</Link></li>     

      {usuariotipo === "estudiante" && (
        <>
        <li><Link to="/miscursos">Mis Cursos</Link></li>
        <li><Link to="/cerrarsesion">Cerrar Sesión</Link></li>

        </>
      )}


      {usuariotipo === "admin" &&(
        <>
        
        <li><Link to="/cursonuevo">Agregar Curso</Link></li>   
        <li><Link to="/servicios">Servicios</Link></li>
        <li><Link to="/cerrarsesion">Cerrar Sesión</Link></li>

        </>
      )}

      {usuariotipo !== "admin" && usuariotipo !== "estudiante" && (
        <>
        <li><Link to="/login">Login</Link></li>
        <li><Link to="/registrarse">Registrar</Link></li>
        </>
      )}

        <li>
          <form onSubmit={handleSearchSubmit}>
            <input
              type="text"
              value={palabra}
              onChange={handleSearchChange}
              placeholder="Buscar cursos..."
            />
            <button type="submit">Buscar</button>
          </form>
        </li>
      </ul>
    </nav>
  );
};

export default NavegadorHome;

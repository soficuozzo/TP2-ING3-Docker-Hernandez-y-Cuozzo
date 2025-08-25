import React, { useState, useEffect } from 'react';
import './miscursos.css';
import axios from 'axios';
import { Link } from 'react-router-dom';

function MisCursos({ searchTerm }) {
  const [cursos, setCursos] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    const usuarioId = localStorage.getItem("usuario_id");
    if (!usuarioId) {
      setError("No se encontró el ID del usuario en el almacenamiento local.");
      return;
    }

    const fetchCursos = async () => {
      setLoading(true);
      setError(null);
      try {
        const response = await axios.get(`http://localhost:8082/usuario/miscursos/${usuarioId}`);
        console.log('Response data:', response.data); // Verifica los datos recibidos desde la API
        if (response.status !== 200) {
          throw new Error('Error en la solicitud');
        }
        const data = response.data;
        if (Array.isArray(data.cursos) && data.cursos.length > 0) {
          setCursos(data.cursos);
        } else {
          setCursos([]);
          setError('No se encontraron cursos.');
        }
      } catch (error) {
        console.error('Error fetching courses:', error);
        setError('Error al obtener los datos');
      } finally {
        setLoading(false);
      }
    };

    fetchCursos();
  }, []);

  return (
    <div className="fullscreen-container">
      <div className="Resultadosfondo">
        {loading ? (
          <img
            style={{ width: '36px' }}
            src="https://raw.githubusercontent.com/Codelessly/FlutterLoadingGIFs/master/packages/circular_progress_indicator_square_large.gif"
            alt="Loading"
          />
        ) : error ? (
          <p>{error}</p>
        ) : (
          <div className="course-list">
            {cursos.length > 0 ? (
              cursos.map(curso => (
                <div key={curso.id} className="course-list-item">
                  <div>
                    <br/>
                      <Link to={`/curso/${curso.id}`} className="curso-link">
                                        <strong>{curso.nombre}</strong>
                                    </Link>
                    <p>{curso.descripcion}</p>
                    <p><b>Categoria:</b> {curso.categoria}</p>
                    <br/>
                  </div>
                  <hr />
                </div>
              ))
            ) : (
              <p>No se encontraron cursos.</p>
            )}
          </div>
        )}
      </div>
    </div>
  );
}

export default MisCursos;

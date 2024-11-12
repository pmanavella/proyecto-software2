import React, { useEffect, useState} from 'react'
import './Details.css';
import { useParams } from 'react-router-dom';
import Cookies from "universal-cookie";

const Cookie = new Cookies();


function goto(path){
    window.location = window.location.origin + path
}

const Details = () => {
    const[course, setCourse] =useState({});
    const { id } = useParams();
    // const navigate = useNavigate();

    // useEffect(() => {
    //   fetch(`http://localhost:8081/course/${id}`)
    //     .then(response => response.json())
    //     .then(data => setCourse(data));
    // }, [id]);
  
    // const handleRegister = () => {
        // InscriptionCourses(); 
    //   const userId = cookies.get('userId'); // Asumiendo que el ID del usuario está almacenado en una cookie
    //   fetch(`http://localhost:8080/register`, {
    //     method: 'POST',
    //     headers: {
    //       'Content-Type': 'application/json',
    //     },
    //     body: JSON.stringify({ courseId: id, userId: userId }),
    //   })
    //     .then(response => {
    //       if (!response.ok) {
    //         throw new Error('Error en la inscripción');
    //       }
    //       return response.json();
    //     })
    //     .then(data => {
    //       alert('Inscripción exitosa');
    //     })
    //     .catch(error => {
    //       alert('Error en la inscripción');
    //       console.error('Error:', error);
    //     });
    // };

    const showCourses= () =>{
        return(
            <div>
                <div className="course-banner-small">detalle del curso</div>
                <div class="course-title">
                    {course.title}
                </div>
                <div class="course-info">
                    <div>
                    <div>
                        <h4>Descripcion:</h4>
                        <p class="course-description">{course.description}</p>
                    </div>
                        <h4>Puntos:</h4>
                        <p class="course-points">{course.points}</p>
                    
                        <h4>categoria: </h4>
                        <p class="course-category">{course.category}</p>
                    </div>
                    <div>
                        <div class="course-image">
                            <img src={course.image_url} alt={course.title}></img>
                        </div>
                    </div>
                    <div>
                        <h4>instructor: </h4>
                        <p class="course-instructor">{course.instructor}</p>
                    </div>
                    <div>
                        <h4>duración: </h4>
                        <p class="course-duration">{course.duration}</p>
                    </div>
                    <div>
                        <h4>capacidad: </h4>
                        <p class="course-capacity">{course.capacity}</p>
                    </div>
                    <div>
                        <div class="course-requirements-container">
                            <h4>requisitos: </h4>
                            <p class="course-requirements">{course.requirements}</p>

                            <button className='green-button' type="submit">inscribirme</button>
                        </div>
                    </div>
                </div>
            </div>
        )
    }

    return (
        <div>
            {showCourses()}
            <button className="white-button" onClick={() => goto("/results")}>⬅️</button>
        </div>
    )
}

export default Details;

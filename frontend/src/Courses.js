import React, { useEffect, useState} from 'react'
import './Courses.css';
import { useParams } from 'react-router-dom';
import Cookies from "universal-cookie";

const Cookie = new Cookies();

/// acá debe ir GetUserCourses
// async function getCourseById(id){
//     return await fetch('http://localhost:8081/course/' + id, {
//     method: 'GET',
//     headers: {
//       'Content-Type': 'application/json'
//     }
// }).then(response => response.json())
// }

function goto(path){
    window.location = window.location.origin + path
}

const Courses = () => {
    const[course, setCourse] =useState({});
    const { id } = useParams();

    // if (!course.id_course) {
    //     getCourseById(Number(id)).then(response => {setCourse(response);})
    // }
    const showCourses= () =>{
        return(
            <div>
                <div className="course-banner-small">Mis cursos</div>
                <div className="course-title">
                    {course.title}
                </div>
                <div className="course-info">
                    <div>
                        <h4>Descripcion:</h4>
                        <p className="course-description">{course.description}</p>
                    </div>
                    <div>
                        <h4>Categoria: </h4>
                        <p className="course-category">{course.category}</p>
                    </div>
                    <div>
                        
                        {/* <p class="course-instructor">{course.instructor}</p> */}
                    </div>
                    <div>
                        <h4>Duración: </h4>
                        {/* <p class="course-duration">{course.duration}</p> */}
                    </div>
                    <div>
                        <h4>Capacidad: </h4>
                        <p className="course-capacity">{course.capacity}</p>
                    </div>
                    <div>
                        <h4>Puntos: </h4>
                        <p className="course-points">{course.points}</p>
                    </div>
                    <div>
                        <h4>Requisitos: </h4>
                        <div className="course-requirements-container">
                            <p className="course-requirements">{course.requirements}</p>
                        </div>
                        <div className="course-image">
                            <img src={course.image_url} alt={course.title}></img>
                        </div>
                    </div>
                </div>
            </div>
        )
    }

    return (
        <div>
            {showCourses()}
            <button className="white-button" onClick={() => goto("/")}>HOME 🏠</button>
            
        </div>
    )
}

export default Courses;

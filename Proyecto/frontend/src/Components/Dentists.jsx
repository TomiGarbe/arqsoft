import React, { useEffect, useState } from 'react'
import Card from "./Card"

const dentistas = "https://jsonplaceholder.typicode.com/users"

const Dentist = () =>  {

    const [dentist, setDentist] = useState([])

    const getDentist = async () => {
        const response = await fetch(dentistas);
        const resolve = await response.json();
        setDentist(resolve)
    }

    useEffect(() => {
        getDentist();
    },[])

    return (
        <div className='gridOut'>
            <h3>Selección de Odontólogos</h3>
            <div className='grid'>
            <div className='dentCard'>
                {
                    dentist.length ? dentist.map((dentist) => <Card key={dentist.id} name={dentist.name} username={dentist.username} id={dentist.id}/>): null
                }
            </div>
            </div>
        </div>
    )
}

export default Dentist
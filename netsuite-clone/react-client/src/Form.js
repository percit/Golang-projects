import React, { useState } from 'react';
import './Form.css';

function Form() {
  const [formData, setFormData] = useState({
    workerName: '',
    date: '',
    hours: ''
  });

  const handleSubmit = (event) => {
    event.preventDefault();

    const urlEncodedData = new URLSearchParams();
    for (const key in formData) {
      urlEncodedData.append(key, formData[key]);
    }
    fetch('http://localhost:8080/api/InsertRecord', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      body: urlEncodedData.toString()
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error('Request failed');
        }
        return response.json();
      })
      .then((data) => {
        console.log(data);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const handleChange = (event) => {
    setFormData({
      ...formData,
      [event.target.name]: event.target.value
    });
  };

  return (
    <div className="form-container">
      <form onSubmit={handleSubmit}>
        <div className="form-field">
          <label htmlFor="workerName">WorkerName:</label>
          <input
            type="text"
            id="workerName"
            name="workerName"
            value={formData.workerName}
            onChange={handleChange}
          />
        </div>
        <div className="form-field">
          <label htmlFor="date">Date:</label>
          <input
            type="text"
            id="date"
            name="date"
            value={formData.date}
            onChange={handleChange}
          />
        </div>
        <div className="form-field">
          <label htmlFor="hours">Hours:</label>
          <input
            type="text"
            id="hours"
            name="hours"
            value={formData.hours}
            onChange={handleChange}
          />
        </div>
        <button type="submit">Submit</button>
      </form>
    </div>
  );
}

export default Form;

import React, { useState } from 'react';
import './Form.css'; // Import the CSS file for styling

function Form() {
  const [field1, setField1] = useState('');
  const [field2, setField2] = useState('');
  const [field3, setField3] = useState('');
  const [submitted, setSubmitted] = useState(false);

  const handleSubmit = (event) => {
    event.preventDefault();
    setSubmitted(true);
    console.log(field1, field2, field3);
  };

  return (
    <div className="form-container">
      <form onSubmit={handleSubmit}>
        <div className="form-field">
          <label htmlFor="field1">Field 1:</label>
          <input
            type="text"
            id="field1"
            value={field1}
            onChange={(event) => setField1(event.target.value)}
          />
        </div>
        <div className="form-field">
          <label htmlFor="field2">Field 2:</label>
          <input
            type="text"
            id="field2"
            value={field2}
            onChange={(event) => setField2(event.target.value)}
          />
        </div>
        <div className="form-field">
          <label htmlFor="field3">Field 3:</label>
          <input
            type="text"
            id="field3"
            value={field3}
            onChange={(event) => setField3(event.target.value)}
          />
        </div>
        <button type="submit">Submit</button>
      </form>
      {submitted && (
        <div>
          <h4>Submitted Values:</h4>
          <p>Field 1: {field1}</p>
          <p>Field 2: {field2}</p>
          <p>Field 3: {field3}</p>
        </div>
      )}
    </div>
  );
}

export default Form;

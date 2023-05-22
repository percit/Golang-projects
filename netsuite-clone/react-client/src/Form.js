import React, { useState } from 'react';
import './Form.css'; // Import the CSS file for styling

function Form() {
  const [field1, setField1] = useState('');
  const [field2, setField2] = useState('');
  const [field3, setField3] = useState('');

  const handleSubmit = (event) => {
    event.preventDefault();
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
    </div>
  );
}

export default Form;

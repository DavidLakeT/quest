const API_BASE_URL = 'http://127.0.0.1:3001/apis/citizen';
export const registerCitizen = async (citizenData) => {
  try {
    const response = await fetch(`${API_BASE_URL}/registerCitizen`, {
      method: 'POST',
      mode: 'no-cors',
      headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Methods': 'POST,GET,OPTIONS',
      },
      body: JSON.stringify(citizenData)
    });

    if (!response.ok) {
      throw new Error('Error al crear ciudadano');
    }

    const createdCitizen = await response.json();
    return createdCitizen;
  } catch (error) {
    throw new Error('Error al crear ciudadano');
  }
};
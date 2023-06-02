const API_BASE_URL = 'http://127.0.0.1:3001/apis/citizen';

export const registerCitizen = async (citizenData) => {
  try {
    const response = await fetch(`${API_BASE_URL}/registerCitizen`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
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

export const transferCitizen = async (transferData) => {
  try {
    const response = await fetch(`${API_BASE_URL}/transferCitizen`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(transferData)
    });

    if (!response.ok) {
      throw new Error('Error al crear ciudadano');
    }

    const transfer = await response.json();
    return transfer;
  } catch (error) {
    throw new Error('Error al crear ciudadano');
  }
};

export const citizenLogin = async (citizenData) => {
  try {
    const response = await fetch(`${API_BASE_URL}/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(citizenData),
      //credentials: 'include'
    });

    if (!response.ok) {
      throw new Error('Error');
    }

    const login = await response

    const cookie = response.headers.get('Set-Cookie');

    document.cookie = cookie;
    return login;
  } catch (error) {
    throw new Error('Error al crear ciudadano');
  }
};

export const citizenDocuments = async (citizenId) => {
  try {
    console.log(citizenId)
  const response = await fetch(`${API_BASE_URL}/getCitizenDocuments/${citizenId}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
  }});

    if (!response.ok) {
      throw new Error('Error al realizar la solicitud');
    }

    const data = await response.json();
    console.log(data.message);
    return data
  } catch (error) {
    console.error('Error al obtener los documentos del ciudadano:', error);
    throw error;
  }
  /*
      .then(response => {
      if (!response.ok) {
        throw new Error('Error al realizar la solicitud');
      }
      return response.json();
    })
    .then(data => {
      // Hacer algo con los datos recibidos en la respuesta
      console.log(data.message);
      return data.message
    })*/
    // Manejar el error en caso de que ocurra
    /*if (!response.ok) {
      throw new Error('Error');
    }

    const documents = await response.json();
    return documents;*/
  /*} catch (error) {
    throw new Error('Error al crear ciudadano');
  }*/
};
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
    const response = await fetch(`${API_BASE_URL}/getCitizenDocuments/${citizenId}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      //body: JSON.stringify(citizenData)
    });

    if (!response.ok) {
      throw new Error('Error');
    }

    const documents = await response.json();
    return documents;
  } catch (error) {
    throw new Error('Error al crear ciudadano');
  }
};
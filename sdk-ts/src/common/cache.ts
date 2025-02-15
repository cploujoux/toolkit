export async function handleControlplaneCache(req:Request):Promise<Response | null> {
  if (req.method != "GET") {
    return null
  }
  if (!req.url.includes('https://api.beamlit.com') && !req.url.includes('https://api.beamlit.dev')) {
    return null
  }
  
  const pathSegments = req.url.split('/')
  if(pathSegments.length > 6) {
    return null
  }
  const objectType = pathSegments[4]
  const name = pathSegments[5]
  const requirePath = `${process.cwd()}/.beamlit/cache/${objectType}/${name}.json`
  
  // Importation dynamique de 'fs'
  let fs;
  try {
    fs = await import('fs');
  } catch {
    return null; // Retourner null si 'fs' n'est pas disponible
  }

  try {
    const cache = fs.readFileSync(requirePath, 'utf8')
    return new Response(cache, {
      status: 200,
      headers: { 'Content-Type': 'application/json' }, // Headers as plain object
    })
  } catch {
    return null
  }
}
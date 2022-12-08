import { useEffect } from 'react';
import { Home } from './components/home';
import { useShortcuts } from './hooks/keypress';
import { addAction } from './handlers/handler';
import { useDrawer } from './hooks/drawer';
import { ProductForm } from './components/product/product-form';
import { Text } from '@chakra-ui/react'

function App() {
  const shortcuts = useShortcuts()
  const drawer = useDrawer()

  useEffect(() => {

    addAction({
      name:"add_product", 
      handler: () => {
        drawer({
          header: <Text as="h1">Novo produto</Text>,
          body: <ProductForm/>
        })
      }
    })
  }, [shortcuts, drawer])

  return <Home />
}

export default App;

# Estruturas de Dados

Este é um guia que explora várias estruturas de dados comuns, suas características e aplicações. Cada estrutura de dados desempenha um papel importante na organização e manipulação eficiente de informações em algoritmos.

## Arrays

Arrays, também conhecidos como vetores, são estruturas de dados que armazenam elementos do mesmo tipo de forma contígua na memória. Os elementos são acessados através de um índice, que representa a posição do elemento no array. Os arrays têm tamanho fixo e são eficientes para acesso direto a elementos, mas não podem ser facilmente redimensionados.

## Listas

As listas são estruturas de dados que armazenam elementos em sequência. Cada elemento é representado por um nó que contém o valor e uma referência para o próximo nó. Diferentemente dos arrays, as listas podem crescer ou diminuir dinamicamente, pois os nós podem ser alocados independentemente na memória. Existem diferentes tipos de listas, como listas ligadas simples, listas duplamente ligadas e listas circulares.

## Pilhas (Stacks)

Pilhas são estruturas de dados baseadas no princípio LIFO (Last-In, First-Out), onde o último elemento inserido é o primeiro a ser removido. As operações principais em uma pilha são "push" para adicionar elementos ao topo da pilha e "pop" para remover o elemento do topo. As pilhas são amplamente utilizadas em algoritmos de recursão, avaliação de expressões matemáticas e gerenciamento de chamadas de função.

## Filas (Queues)

Filas são estruturas de dados baseadas no princípio FIFO (First-In, First-Out), onde o primeiro elemento inserido é o primeiro a ser removido. As operações principais em uma fila são "enqueue" para adicionar elementos ao final da fila e "dequeue" para remover o elemento do início da fila. As filas são úteis em algoritmos de agendamento, processamento de tarefas em ordem e gerenciamento de recursos compartilhados.

## Árvores (Trees)

Árvores são estruturas de dados hierárquicas compostas por nós conectados por arestas. Cada nó pode ter zero ou mais nós filhos, e um nó sem filhos é chamado de folha. A parte superior da árvore é chamada de raiz. Existem vários tipos de árvores, como árvores binárias, árvores de busca binária, árvores balanceadas e árvores de segmentos. As árvores são amplamente utilizadas em estruturas de diretórios, algoritmos de pesquisa e otimização.

## Grafos (Graphs)

Grafos são estruturas de dados compostas por vértices (ou nós) e arestas (ou conexões) que representam as relações entre os vértices. Os grafos podem ser direcionados, onde as arestas têm uma direção, ou não direcionados. Eles são usados para modelar uma variedade de problemas do mundo real, como redes sociais, rotas de transporte, jogos e algoritmos de busca.

## Tabelas Hash (Hash Tables)

Tabelas hash são estruturas de dados que mapeiam chaves para valores. Elas

 usam uma função hash para calcular um índice onde o valor correspondente é armazenado. As tabelas hash permitem acesso rápido aos valores com base em suas chaves e são eficientes para inserção, exclusão e busca médias. Elas são amplamente utilizadas em bancos de dados, caches de memória, indexação de dados e algoritmos de pesquisa eficientes.

## Conjuntos (Sets)

Conjuntos são estruturas de dados que armazenam elementos únicos sem uma ordem específica. Eles suportam operações como adição de elementos, remoção de elementos e verificação de pertencimento. Conjuntos são úteis em problemas que exigem a verificação de duplicatas ou a operação de união, interseção e diferença entre conjuntos.

## Mapas (Maps)

Mapas, também conhecidos como dicionários, são estruturas de dados que mapeiam chaves para valores. Eles associam um valor a uma chave específica, permitindo uma recuperação eficiente dos valores com base nas chaves. Os mapas são úteis em problemas que envolvem associações entre dados, como contagem de ocorrências, indexação e armazenamento de metadados.

## Listas Ligadas (Linked Lists)

Listas ligadas são estruturas de dados compostas por nós, onde cada nó contém um valor e uma referência para o próximo nó da lista. A última referência da lista é nula, indicando o final da lista. As listas ligadas são eficientes para inserção e remoção de elementos em qualquer posição, mas o acesso direto aos elementos é mais lento em comparação com os arrays. Existem diferentes tipos de listas ligadas, como listas simplesmente ligadas, listas duplamente ligadas e listas circulares.

## Filas de Prioridade (Priority Queues)

Filas de prioridade são estruturas de dados que armazenam elementos com base em suas prioridades. Cada elemento possui uma prioridade associada e o elemento com maior prioridade é sempre o próximo a ser removido. As filas de prioridade são úteis em algoritmos que envolvem agendamento, roteamento, programação de tarefas e simulações.

## Deques (Double-ended Queues)

Deques, também conhecidos como filas de dupla extremidade, são estruturas de dados que permitem a inserção e remoção de elementos em ambos os extremos. Eles suportam operações como inserção no início, inserção no final, remoção do início e remoção do final. Deques são úteis em algoritmos que exigem acesso eficiente tanto no início quanto no final da estrutura de dados.

## Heaps

Heaps são estruturas de dados que representam árvores binárias parcialmente ordenadas. Em um heap, o elemento pai tem uma prioridade maior (ou menor) do que seus filhos. Os heaps são frequentemente usados em algoritmos de classificação, algoritmos de busca e filas de prioridade.

## Tries

Tries, também conhecidos como árvores de prefixos, são estruturas de dados eficientes para armazenar e pesquisar palavras ou sequências de caracteres. Cada nó representa um caractere e os nós formam uma estrutura de árvore onde cada caminho da raiz para um nó folha representa uma palavra completa. Tries são comumente usados em aplicações de pesquisa de texto, autocorreção e dicionários.

Este guia oferece uma visão geral das estruturas de dados mais comuns, mas há muitas outras estruturas e variações a serem exploradas. Cada estrutura de dados tem suas próprias vantagens e desvantagens, e a escolha da estrutura correta depende do problema específico que está sendo resolvido.
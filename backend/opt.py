l1 = [['SHORT', 'FILM'], ['LEASE', 'BACK'], ['SHELF', 'LIFE'], ['HOLDS', 'FAST']]
l2 = [['BOAT', 'DECK'], ['FAST', 'FOOD'], ['FILM', 'PROP'], ['CHOW', 'LINE']]
l3 = [['FOOD', 'DRIVE'], ['PROP', 'PLANE'], ['GOAL', 'LINES'], ['WRAP', 'PARTY']]
l4 = [['FOOD', 'PLATE'], ['PLANE', 'SKY'], ['GOAL', 'LINES'], ['WRAP', 'PARTY']]

final = []

def diff(li1, li2):
    return (list(list(set(li1)-set(li2)) + list(set(li2)-set(li1))))

def compute_current(i, j):
    current = []
    inters = list(set(i).intersection(j))        
    if inters:
        l = diff(i, j)
        for index, item in enumerate(l):
            if index == 1:
               current.append(inters[0])
            if item:
               current.append(item)
    return current 

def unite_lists(li1, li2):
    for i in li2:
        if i not in li1:
            li1.append(i)
    return li1

def shorten(l1, l2, l3, final):
    for i in l1:
        for j in l2:
            inters = list(set(i).intersection(j))
            current = compute_current(i, j)
 
            if current:    
                for k in l3:
     
                    r = compute_current(j, k) 
                    if r:
                        res = [current[0]]
                        res = res + r
                        if res:
                            final.append(res)
    
    return final

def shorten_pair(l1, l2, final):
    fn = []
    for i in l1:
        for j in l2:
            inters = list(set(i).intersection(j))
            current = compute_current(i, j)
            if current:
                final.append(current)
    #for i in final:
    #    for j in final:
    #        result = unite_lists(i, j)
    #        print(f"RESULT WAS {result}")
    #        if result:
    #            fn.append(result)
    #
    for i in final:
        for j in final:
            if diff(i, j):
                result = unite_lists(i, j)
                if result:
                    print(f"WANT TO APPEND {result} to {fn}")
                    fn.append(result)
    print(f"RETURN {fn}")
    return final


final = shorten_pair(l1, l2, [])
print(final)

final = shorten_pair(l2, l3, final)
print(final)

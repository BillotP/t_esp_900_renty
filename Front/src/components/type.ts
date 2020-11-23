export interface Ticket {
    name: string;
    needs: string;
    priority: number;
    responsible: string;
}

export interface DataTicketReceived {
    items: Ticket[];
}
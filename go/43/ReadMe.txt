Compact Double Link List.

This Package provides a double Linked List Functionality.

The List uses a compact Data Model. This means that, as opposed to the
Golang's built-in 'list' Package, this Package does not store the Pointer
to the Owner-List in each List Item. This greatly reduces Memory Load for
Lists that have simple Data Model and big Size, but, as a Drawback, we
have to do thorough Calculations each Time we want to delete an Item from
the List (when the deleted Item is neither Head, nor Tail of the List).
However, it is important to say that Tail and Head Items are deleted in a
very fast Manner.

The Package provides extended Functionality comparing with the Golang's
built-in 'list' Package.
